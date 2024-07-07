import React, { useState } from 'react';
import { loginAdmin, loginStudent } from './api';
import { useAuth } from './AuthContext';
import { useNavigate } from 'react-router-dom';
import toast, { Toaster } from 'react-hot-toast';

const LoginStudent = () => {
  const [usn, setUsn] = useState('');
  const [password, setPassword] = useState('');
  const userType = 'student';
  const { login } = useAuth();
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const data = await loginStudent(usn, password);
      console.log(data)
      login(data.token, userType,usn);
      navigate('/session');

    } catch (error) {
      console.log(error.toString())
      toast.error(error.toString())
    }
  };

  return (
    <div>
      <Toaster position='top-right'/>
    <form onSubmit={handleSubmit}>
      <input
        type="usn"
        value={usn}
        onChange={(e) => setUsn(e.target.value)}
        placeholder="USN"
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
    </div>
  );
};

export default LoginStudent;