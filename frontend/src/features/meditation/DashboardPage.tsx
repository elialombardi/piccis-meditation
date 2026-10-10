import {
  AppBar, Avatar, Box, Card, CardActionArea, CardContent, Container,
  Grid, IconButton, Toolbar, Typography,
} from '@mui/material';
import * as Icons from '@mui/icons-material';
import LogoutIcon from '@mui/icons-material/Logout';
import HistoryIcon from '@mui/icons-material/History';
import { useDispatch } from 'react-redux';
import { useNavigate } from 'react-router-dom';
import { logout } from '../auth/authSlice';
import { useListTypesQuery } from './meditationApi';

export default function DashboardPage() {
  const { data: types = [], isLoading } = useListTypesQuery();
  const dispatch = useDispatch();
  const navigate = useNavigate();

  const signOut = () => {
    dispatch(logout());
    navigate('/login');
  };

  return (
    <Box>
      <AppBar position="static" color="transparent" elevation={0}>
        <Toolbar>
          <Typography variant="h6" sx={{ flexGrow: 1 }}>Picci's Meditation</Typography>
          <IconButton onClick={() => navigate('/habits')} title="Habits"><HistoryIcon /></IconButton>
          <IconButton onClick={signOut} title="Sign out"><LogoutIcon /></IconButton>
        </Toolbar>
      </AppBar>

      <Container sx={{ mt: 4 }}>
        <Typography variant="h4" gutterBottom>Choose a practice</Typography>
        <Typography variant="body1" color="text.secondary" sx={{ mb: 3 }}>
          Pick a type and a duration. We'll guide you with voice, step by step.
        </Typography>

        {isLoading && <Typography>Loading…</Typography>}

        <Grid container spacing={2}>
          {types.map((t) => {
            const Icon = (Icons as any)[t.icon] ?? Icons.SelfImprovement;
            return (
              <Grid item xs={12} sm={6} md={4} key={t.id}>
                <Card>
                  <CardActionArea onClick={() => navigate(`/meditate/${t.slug}`)}>
                    <CardContent>
                      <Avatar sx={{ bgcolor: 'primary.main', mb: 2 }}>
                        <Icon />
                      </Avatar>
                      <Typography variant="h6">{t.name}</Typography>
                      <Typography variant="body2" color="text.secondary">
                        {t.description}
                      </Typography>
                    </CardContent>
                  </CardActionArea>
                </Card>
              </Grid>
            );
          })}
        </Grid>
      </Container>
    </Box>
  );
}