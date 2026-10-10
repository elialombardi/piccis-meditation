import { useState } from 'react';
import { Box, Button, Paper, TextField, Typography, Link } from '@mui/material';
import { Link as RouterLink, useNavigate } from 'react-router-dom';
import { useDispatch } from 'react-redux';
import { useRegisterMutation } from './authApi';
import { setCredentials } from './authSlice';

export default function RegisterPage() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [register, { isLoading, error }] = useRegisterMutation();
  const dispatch = useDispatch();
  const navigate = useNavigate();

  const onSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const res = await register({ email, password }).unwrap();
      dispatch(setCredentials(res));
      navigate('/');
    } catch {}
  };

  return (
    <Box display="flex" justifyContent="center" alignItems="center" minHeight="100vh">
      <Paper sx={{ p: 4, width: 360 }}>
        <Typography variant="h5" gutterBottom>Create account</Typography>
        <form onSubmit={onSubmit}>
          <TextField fullWidth label="Email" margin="normal" value={email} onChange={(e) => setEmail(e.target.value)} />
          <TextField fullWidth label="Password" type="password" margin="normal" value={password} onChange={(e) => setPassword(e.target.value)} />
          {error && <Typography color="error" variant="body2">Registration failed</Typography>}
          <Button type="submit" fullWidth variant="contained" sx={{ mt: 2 }} disabled={isLoading}>Register</Button>
        </form>
        <Typography variant="body2" sx={{ mt: 2 }}>
          Already have an account? <Link component={RouterLink} to="/login">Login</Link>
        </Typography>
      </Paper>
    </Box>
  );
}