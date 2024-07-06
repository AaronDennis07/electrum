import React, { useState } from 'react';
import { loginAdmin } from './api';
import { useAuth } from './AuthContext';
import { useNavigate } from 'react-router-dom';

const LoginAdmin = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const userType = 'admin';
  const { login } = useAuth();
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const data = await loginAdmin(email, password);
      login(data.token, userType);
      navigate(userType === 'admin' ? '/admin/session' : '/session');

    } catch (error) {
      console.error('Login failed:', error);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        placeholder="Email"
        required
      />
      <input
        type="password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        placeholder="Password"
        required
      />
      
      <button type="submit">Login</button>
    </form>
  );
};

export default LoginAdmin;