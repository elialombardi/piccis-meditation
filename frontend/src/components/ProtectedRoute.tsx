import { Navigate } from 'react-router-dom';
import { useSelector } from 'react-redux';
import { ReactNode } from 'react';

export default function ProtectedRoute({ children }: { children: ReactNode }) {
  const token = useSelector((s: any) => s.auth.token);
  return token ? <>{children}</> : <Navigate to="/login" replace />;
}