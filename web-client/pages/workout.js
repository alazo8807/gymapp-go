
import { useEffect, useState } from 'react';
import { getWorkouts } from '../utils/api';
import { DataGrid } from '@mui/x-data-grid';

const columns = [
  { field: 'description', headerName: 'Description', width: 130 },
  { field: 'created_at', headerName: 'Date', width: 130 },
];

const WorkoutsPage = () => {
  const [workouts, setWorkouts] = useState([]);

  useEffect(() => {
    const fetchWorkouts = async () => {
      try {
        const { data } = await getWorkouts();
        setWorkouts(data);
      } catch (error) {
        console.error('Error fetching workouts:', error);
      }
    };

    fetchWorkouts();
  }, []);

  return (
    <div style={{ height: 400, width: '100%' }}>
      <DataGrid
        rows={workouts}
        columns={columns}
        initialState={{
          pagination: {
            paginationModel: { page: 0, pageSize: 5 },
          },
        }}
        pageSizeOptions={[5, 10]}
        checkboxSelection
      />
    </div>
  );
};

export default WorkoutsPage;
