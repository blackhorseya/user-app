import {Route, Routes} from 'react-router-dom';
import Home from '@/pages';
import Profile from '@/pages/profile';
import {ProtectedRoute} from '@/store';

function App() {
  return (
      <Routes>
        <Route path="/" element={<Home/>}>
          <Route
              path="/profile"
              element={
                <ProtectedRoute>
                  <Profile/>
                </ProtectedRoute>
              }
          />
        </Route>
      </Routes>
  );
}

export default App;