
import CreateSessionForm from './CreateSession.jsx';
// import { BrowserRouter as Router, Route, Switch, Link } from 'react-router-dom';
import {  BrowserRouter, Route, Router, Routes } from "react-router-dom";
import { UserProvider, Login } from './UserContext.jsx';
import { Navigate } from 'react-router-dom';
import EnrollmentPeriodCourses from './Enrollment.jsx';
import SessionListPage from './SessionList.jsx';
import DownloadSession from './DownloadSession.jsx';
import AdminSessionPage from './AdminSession.jsx';
import UploadStudent from './UploadStudent.jsx';
import AdminSessionDashboard from './SessionDetails.jsx';
import { AuthProvider } from './AuthContext.jsx';
import LoginAdmin from './LoginAdmin.jsx';
import PrivateRoute from './PrivateRoute.jsx';


export function App() {
  return (
   
    <AuthProvider>
    <BrowserRouter> 
  
      <Routes>
      <Route path="/admin/login" element={<LoginAdmin/>} />
      <Route path="/admin/create" element={<PrivateRoute allowedUserType="admin"><CreateSessionForm/></PrivateRoute>} />
      <Route path="/admin/download" element={<PrivateRoute allowedUserType="admin"><DownloadSession/></PrivateRoute>} />
      <Route path="/admin/session" element={<PrivateRoute allowedUserType="admin"><AdminSessionPage/></PrivateRoute>} />
      <Route path="/admin/upload" element={<PrivateRoute allowedUserType="admin"><UploadStudent/></PrivateRoute>} />
      <Route path="/admin/session/:sessionName" element={<PrivateRoute allowedUserType="admin"><AdminSessionDashboard/></PrivateRoute>} />

      {/* <Route path="/enroll" element={<EnrollmentPeriodCourses />} /> */}
      {/* <Route path="/create" element={<CreateSessionForm />} /> */}
      <Route path="/enroll/:sessionName" element={<EnrollmentPeriodCourses/>} />
      <Route path='/session' element={<SessionListPage/>} />
      {/* <Route path="/enroll/:sessionName" element={<EnrollmentPeriodCourses />} /> */}
      {/* <Route path="/download" element={<DownloadSession />} /> */}
      {/* <Route path="/admin/session" element={<AdminSessionPage />} /> */}
      {/* <Route path="/admin/upload" element={<UploadStudent />} /> */}

// In your Routes element
{/* <Route path="/admin/session/:sessionName" element={<AdminSessionDashboard />} /> */}
      {/* <Route path="/*" element={<Navigate to="/login" replace />} /> */}
      <Route path="/" element={<Navigate to="/login" replace />} />
      </Routes>
           
      </BrowserRouter>
    </AuthProvider>

  );
}
export default App;