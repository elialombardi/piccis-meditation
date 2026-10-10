import { createTheme } from '@mui/material/styles';

export const theme = createTheme({
  palette: {
    mode: 'dark',
    primary:   { main: '#8ab4f8' },
    secondary: { main: '#c58af9' },
    background: { default: '#0b0f14', paper: '#141a22' },
    text: { primary: '#e8eaed', secondary: '#9aa0a6' },
  },
  shape: { borderRadius: 14 },
  typography: {
    fontFamily: '"Inter", system-ui, -apple-system, sans-serif',
    h4: { fontWeight: 300, letterSpacing: '0.02em' },
  },
});