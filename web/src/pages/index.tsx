import {useEffect, useMemo} from 'react';
import {Outlet, useNavigate, useSearchParams} from 'react-router-dom';
import {observer} from 'mobx-react-lite';
import rootStore from '@/store';
import Navbar from '@/layouts/navbar';

const Home = () => {
  const [searchParams] = useSearchParams();
  const token = useMemo(() => searchParams.get('token'), [searchParams]);
  const navigate = useNavigate();

  useEffect(() => {
    if (token) {
      rootStore.login(token);
      navigate('/');
    }
  }, [token]);

  return (
      <div className="flex flex-col full-screen">
        <Navbar/>
        <div className="flex-grow">
          <Outlet/>
        </div>
      </div>
  );
};

export default observer(Home);