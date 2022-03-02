import {PropsWithChildren} from 'react';
import {Navigate} from 'react-router-dom';
import {observer} from 'mobx-react-lite';
import rootStore from '@/store';

interface IProps {
  redirectTo?: string;
}

const ProtectedRoute = ({
                          redirectTo = '/',
                          children,
                        }: PropsWithChildren<IProps>) => {
  if (!rootStore.isAuth) return <Navigate to={redirectTo} replace/>;
  return <>{children}</>;
};

export default observer(ProtectedRoute);