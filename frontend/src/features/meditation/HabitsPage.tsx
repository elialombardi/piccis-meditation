import {
  AppBar, Box, Card, CardContent, Container, Grid, IconButton,
  List, ListItem, ListItemText, Toolbar, Typography,
} from '@mui/material';
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import { useNavigate } from 'react-router-dom';
import { useGetStatsQuery, useListSessionsQuery } from './meditationApi';

export default function HabitsPage() {
  const navigate = useNavigate();
  const { data: stats, isLoading: statsLoading } = useGetStatsQuery();
  const { data: sessions = [] } = useListSessionsQuery();

  const maxCount = Math.max(1, ...(stats?.last30Days ?? []).map((d) => d.count));

  return (
    <Box>
      <AppBar position="static" color="transparent" elevation={0}>
        <Toolbar>
          <IconButton edge="start" onClick={() => navigate('/')}><ArrowBackIcon /></IconButton>
          <Typography variant="h6" sx={{ ml: 1 }}>Your habits</Typography>
        </Toolbar>
      </AppBar>

      <Container sx={{ mt: 4 }}>
        {statsLoading || !stats ? (
          <Typography>Loading…</Typography>
        ) : (
          <>
            <Grid container spacing={2} sx={{ mb: 4 }}>
              <StatCard label="Total sessions" value={stats.totalSessions} />
              <StatCard label="Minutes practiced" value={stats.totalMinutes} />
              <StatCard label="Current streak" value={`${stats.currentStreakDays}d`} />
              <StatCard label="Longest streak" value={`${stats.longestStreakDays}d`} />
            </Grid>

            <Typography variant="h6" gutterBottom>Last 30 days</Typography>
            <Box sx={{ display: 'flex', gap: 0.5, alignItems: 'flex-end', height: 80, mb: 4 }}>
              {stats.last30Days.map((d) => (
                <Box
                  key={d.date}
                  title={`${d.date}: ${d.count}`}
                  sx={{
                    flex: 1,
                    height: `${(d.count / maxCount) * 100}%`,
                    minHeight: 2,
                    bgcolor: d.count > 0 ? 'primary.main' : 'divider',
                    borderRadius: 1,
                  }}
                />
              ))}
            </Box>

            <Typography variant="h6" gutterBottom>Recent sessions</Typography>
            <List>
              {sessions.map((s) => (
                <ListItem key={s.id}>
                  <ListItemText
                    primary={s.typeName}
                    secondary={`${Math.round(s.durationSeconds / 60)} min · ${new Date(s.completedAt).toLocaleString()}`}
                  />
                </ListItem>
              ))}
              {sessions.length === 0 && (
                <Typography color="text.secondary">No sessions yet — start your first one!</Typography>
              )}
            </List>
          </>
        )}
      </Container>
    </Box>
  );
}

function StatCard({ label, value }: { label: string; value: number | string }) {
  return (
    <Grid item xs={6} md={3}>
      <Card>
        <CardContent>
          <Typography variant="overline" color="text.secondary">{label}</Typography>
          <Typography variant="h4">{value}</Typography>
        </CardContent>
      </Card>
    </Grid>
  );
}