import { Navigate, useLocation } from "react-router-dom";
import { useAuth } from "./AuthContext";

const PrivateRoute = ({ children, allowedUserType }) => {
  const { user } = useAuth();
  const location = useLocation();
  console.log(user);
  if (!user) {
    return <Navigate to="/login" state={{ from: location }} replace />;
  }

  if (user.userType !== allowedUserType) {
    return <Navigate to="/" replace />;
  }

  return children;
};

export default PrivateRoute;
