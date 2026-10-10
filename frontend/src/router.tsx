import { createBrowserRouter } from 'react-router-dom';
import LoginPage from './features/auth/LoginPage';
import RegisterPage from './features/auth/RegisterPage';
import DashboardPage from './features/meditation/DashboardPage';
import MeditationSessionPage from './features/meditation/MeditationSessionPage';
import HabitsPage from './features/meditation/HabitsPage';
import ProtectedRoute from './components/ProtectedRoute';

export const router = createBrowserRouter([
  { path: '/login', element: <LoginPage /> },
  { path: '/register', element: <RegisterPage /> },
  { path: '/', element: <ProtectedRoute><DashboardPage /></ProtectedRoute> },
  { path: '/meditate/:slug', element: <ProtectedRoute><MeditationSessionPage /></ProtectedRoute> },
  { path: '/habits', element: <ProtectedRoute><HabitsPage /></ProtectedRoute> },
]);