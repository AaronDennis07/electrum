const API_URL = "http://localhost:8000/auth"; // Replace with your API URL

export const loginStudent = async (usn, password) => {
  console.log(usn, password);
  const response = await fetch(`${API_URL}/student/login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ usn, password }),
  });

  if (!response.ok) {
    const error = await response.json();
    let err = "";
    if (error.error) err = error.error;
    else err = "Something went wrong";

    throw new Error(err);
  }

  return response.json();
};
export const registerStudent = async (usn, email, password) => {
  console.log(usn, password);
  const response = await fetch(`${API_URL}/student/register`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ usn, email, password }),
  });

  if (!response.ok) {
    const error = await response.json();
    let err = "";
    if (error.error) err = error.error;
    else err = "Something went wrong";

    throw new Error(err);
  }

  return response.json();
};
export const loginAdmin = async (email, password) => {
  const response = await fetch(`${API_URL}/admin/login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ email, password }),
  });

  if (!response.ok) {
    const error = await response.json();
    let err = "";
    if (error.error) err = error.error;
    else err = "Something went wrong";

    throw new Error(err);
  }

  return response.json();
};

export const fetchWithAuth = async (url, options = {}) => {
  const token = localStorage.getItem("token");
  const headers = {
    ...options.headers,
    Authorization: `Bearer ${token}`,
  };

  const response = await fetch(url, { ...options, headers });

  if (response.status === 401) {
    localStorage.removeItem("token");
    localStorage.removeItem("userType");
    window.location.href = "/login";
    throw new Error("Unauthorized");
  }

  return response;
};
