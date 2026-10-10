import { useEffect, useMemo, useRef, useState } from 'react';
import {
  Box, Button, CircularProgress, Container, IconButton, Paper,
  Stack, Typography,
} from '@mui/material';
import CloseIcon from '@mui/icons-material/Close';
import PauseIcon from '@mui/icons-material/Pause';
import PlayArrowIcon from '@mui/icons-material/PlayArrow';
import { useNavigate, useParams } from 'react-router-dom';
import { useCreateSessionMutation, useListTypesQuery } from './meditationApi';
import { useSpeech } from './useSpeech';

type Phase = { label: string; seconds: number; speak: string };

export default function MeditationSessionPage() {
  const { slug } = useParams<{ slug: string }>();
  const navigate = useNavigate();
  const { data: types = [] } = useListTypesQuery();
  const [createSession] = useCreateSessionMutation();
  const { speak, cancel } = useSpeech();

  const type = useMemo(() => types.find((t) => t.slug === slug), [types, slug]);

  // 0 = setup, 1 = running, 2 = done
  const [step, setStep] = useState(0);
  const [targetMinutes, setTargetMinutes] = useState(5);
  const [elapsed, setElapsed] = useState(0);   // seconds since start
  const [paused, setPaused] = useState(false);
  const [phaseIdx, setPhaseIdx] = useState(0);
  const [phaseElapsed, setPhaseElapsed] = useState(0);
  const tickRef = useRef<number | null>(null);

  // Flatten script into a loop of phases bounded by the user's target duration.
  const phases: Phase[] = type?.script?.phases ?? [];
  const repeat = type?.script?.repeat ?? 1;
  const targetSeconds = targetMinutes * 60;

  // Reset when the type changes
  useEffect(() => {
    setStep(0); setElapsed(0); setPhaseIdx(0); setPhaseElapsed(0); setPaused(false);
  }, [slug]);

  // Main timer
  useEffect(() => {
    if (step !== 1 || paused || phases.length === 0) return;
    tickRef.current = window.setInterval(() => {
      setElapsed((e) => {
        const next = e + 1;
        if (next >= targetSeconds) {
          finish(next);
          return targetSeconds;
        }
        return next;
      });
      setPhaseElapsed((pe) => {
        const cur = phases[phaseIdx % phases.length];
        if (pe + 1 >= cur.seconds) {
          const nextIdx = phaseIdx + 1;
          if (nextIdx >= phases.length * repeat) {
            // Let the target-duration timer handle completion; loop phase index.
            setPhaseIdx(nextIdx % phases.length);
          } else {
            setPhaseIdx(nextIdx);
          }
          speak(phases[nextIdx % phases.length].speak);
          return 0;
        }
        return pe + 1;
      });
    }, 1000);
    return () => {
      if (tickRef.current) window.clearInterval(tickRef.current);
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [step, paused, phaseIdx, phases.length]);

  // Speak first phase when starting
  useEffect(() => {
    if (step === 1 && phases.length > 0) speak(phases[0].speak);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [step]);

  const finish = async (seconds: number) => {
    if (tickRef.current) window.clearInterval(tickRef.current);
    cancel();
    setStep(2);
    if (type) {
      try {
        await createSession({
          meditationTypeId: type.id,
          durationSeconds: seconds,
        }).unwrap();
      } catch { /* swallow */ }
    }
  };

  const start = () => { setElapsed(0); setPhaseIdx(0); setPhaseElapsed(0); setPaused(false); setStep(1); };
  const stopEarly = () => finish(elapsed);

  if (!type) {
    return (
      <Container sx={{ mt: 6 }}>
        <Typography>Practice not found.</Typography>
        <Button onClick={() => navigate('/')}>Back</Button>
      </Container>
    );
  }

  const currentPhase = phases[phaseIdx % Math.max(phases.length, 1)];
  const phaseProgress = currentPhase ? (phaseElapsed / currentPhase.seconds) * 100 : 0;
  const overallProgress = (elapsed / targetSeconds) * 100;

  return (
    <Container maxWidth="sm" sx={{ py: 6 }}>
      <Stack direction="row" alignItems="center" sx={{ mb: 3 }}>
        <IconButton onClick={() => { cancel(); navigate('/'); }}><CloseIcon /></IconButton>
        <Typography variant="h6" sx={{ ml: 1, flexGrow: 1 }}>{type.name}</Typography>
      </Stack>

      {step === 0 && (
        <Paper sx={{ p: 4 }}>
          <Typography variant="h5" gutterBottom>Set your duration</Typography>
          <Typography color="text.secondary" sx={{ mb: 3 }}>{type.description}</Typography>
          <Stack direction="row" spacing={1} sx={{ mb: 3 }}>
            {[3, 5, 10, 15, 20].map((m) => (
              <Button
                key={m}
                variant={targetMinutes === m ? 'contained' : 'outlined'}
                onClick={() => setTargetMinutes(m)}
              >{m} min</Button>
            ))}
          </Stack>
          <Button variant="contained" size="large" fullWidth onClick={start}>
            Begin
          </Button>
        </Paper>
      )}

      {step === 1 && (
        <Paper sx={{ p: 4, textAlign: 'center' }}>
          <Box sx={{ position: 'relative', display: 'inline-flex', mb: 3 }}>
            <CircularProgress
              variant="determinate"
              value={phaseProgress}
              size={220}
              thickness={2}
            />
            <Box sx={{
              position: 'absolute', inset: 0, display: 'flex',
              flexDirection: 'column', alignItems: 'center', justifyContent: 'center',
            }}>
              <Typography variant="h4">{currentPhase?.label}</Typography>
              <Typography variant="h6" color="text.secondary">
                {currentPhase ? currentPhase.seconds - phaseElapsed : 0}s
              </Typography>
            </Box>
          </Box>

          <Typography variant="body2" color="text.secondary" sx={{ mb: 1 }}>
            {Math.floor(elapsed / 60)}:{String(elapsed % 60).padStart(2, '0')} / {targetMinutes}:00
          </Typography>
          <Box sx={{ height: 4, bgcolor: 'divider', borderRadius: 2, mb: 3 }}>
            <Box sx={{
              height: '100%', width: `${overallProgress}%`,
              bgcolor: 'primary.main', borderRadius: 2, transition: 'width 1s linear',
            }} />
          </Box>

          <Stack direction="row" spacing={2} justifyContent="center">
            <IconButton onClick={() => setPaused((p) => !p)}>
              {paused ? <PlayArrowIcon /> : <PauseIcon />}
            </IconButton>
            <Button onClick={stopEarly} variant="outlined">Finish early</Button>
          </Stack>
        </Paper>
      )}

      {step === 2 && (
        <Paper sx={{ p: 4, textAlign: 'center' }}>
          <Typography variant="h4" gutterBottom>Well done 🌿</Typography>
          <Typography color="text.secondary" sx={{ mb: 3 }}>
            You meditated for {Math.floor(elapsed / 60)} min {elapsed % 60} sec.
          </Typography>
          <Stack direction="row" spacing={2} justifyContent="center">
            <Button variant="outlined" onClick={() => navigate('/')}>Home</Button>
            <Button variant="contained" onClick={() => navigate('/habits')}>See habits</Button>
          </Stack>
        </Paper>
      )}
    </Container>
  );
}