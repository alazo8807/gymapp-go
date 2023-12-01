const API_BASE_URL = 'http://localhost'; // Replace with your API base URL

export const getWorkouts = async () => {
  const response = await fetch(`${API_BASE_URL}/workout`);
  const data = await response.json();
  return data;
};
