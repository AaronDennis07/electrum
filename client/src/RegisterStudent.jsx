import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { registerStudent } from "./api";
import toast, { Toaster } from "react-hot-toast";

const RegisterStudent = () => {
  const [usn, setUSN] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await registerStudent(usn, email, password);
      toast.success("Registration successful");
      navigate("/login");
    } catch (error) {
      console.log(error.toString());
      toast.error(error.toString());
    }
  };

  return (
    <div>
      <Toaster position="top-right" />
      <h2>Register</h2>
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="usn">USN:</label>
          <input
            type="text"
            id="usn"
            value={usn}
            onChange={(e) => setUSN(e.target.value)}
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

        <button type="submit">Register</button>
      </form>
    </div>
  );
};

export default RegisterStudent;
