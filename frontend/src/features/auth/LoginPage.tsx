import { useState } from 'react';
import { Box, Button, Paper, TextField, Typography, Link } from '@mui/material';
import { Link as RouterLink, useNavigate } from 'react-router-dom';
import { useDispatch } from 'react-redux';
import { useLoginMutation } from './authApi';
import { setCredentials } from './authSlice';

export default function LoginPage() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [login, { isLoading, error }] = useLoginMutation();
  const dispatch = useDispatch();
  const navigate = useNavigate();

  const onSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const res = await login({ email, password }).unwrap();
      dispatch(setCredentials(res));
      navigate('/');
    } catch {}
  };

  return (
    <Box display="flex" justifyContent="center" alignItems="center" minHeight="100vh">
      <Paper sx={{ p: 4, width: 360 }}>
        <Typography variant="h5" gutterBottom>Sign in</Typography>
        <form onSubmit={onSubmit}>
          <TextField fullWidth label="Email" margin="normal" value={email} onChange={(e) => setEmail(e.target.value)} />
          <TextField fullWidth label="Password" type="password" margin="normal" value={password} onChange={(e) => setPassword(e.target.value)} />
          {error && <Typography color="error" variant="body2">Login failed</Typography>}
          <Button type="submit" fullWidth variant="contained" sx={{ mt: 2 }} disabled={isLoading}>Login</Button>
        </form>
        <Typography variant="body2" sx={{ mt: 2 }}>
          No account? <Link component={RouterLink} to="/register">Register</Link>
        </Typography>
      </Paper>
    </Box>
  );
}