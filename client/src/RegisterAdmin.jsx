import React, { useState } from 'react';
import { useHistory } from 'react-router-dom';
import { registerUser } from './api';

const RegisterAdmin = () => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [userType, setUserType] = useState('admin');
  const [error, setError] = useState('');
  const history = useHistory();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    try {
      await registerUser(name, email, password, userType);
      history.push('/login');
    } catch (error) {
      setError(error.message);
    }
  };

  return (
    <div>
      <h2>Register </h2>
      {error && <p style={{ color: 'red' }}>{error}</p>}
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="name">Name:</label>
          <input
            type="text"
            id="name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
          />
        </div>
        <div>
          <label htmlFor="email">Email:</label>
          <input
            type="email"
            id="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </div>
        <div>
          <label htmlFor="password">Password:</label>
          <input
            type="password"
            id="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>
        <div>
          <label htmlFor="userType">User Type:</label>
          <select
            id="userType"
            value={userType}
            onChange={(e) => setUserType(e.target.value)}
          >
            <option value="student">Student</option>
            <option value="admin">Admin</option>
          </select>
        </div>
        <button type="submit">Register</button>
      </form>
    </div>
  );
};

export default RegisterAdmin;