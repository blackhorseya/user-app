import {makeAutoObservable, reaction, runInAction} from 'mobx';
import {enableLogging} from 'mobx-logger';
import {getProfile} from '@/utils/api';
import ProtectedRoute from './ProtectedRoute';
import {User} from '@/types/user';

enableLogging({
  predicate: () => Boolean(window.navigator.userAgent),
  action: true,
  transaction: true,
  reaction: true,
  compute: true,
});

class RootStore {
  token;
  user?: User;

  constructor() {
    makeAutoObservable(this);

    const token = localStorage.getItem('token');
    if (token) {
      this.token = token;
      this.fetchUser();
    }
  }

  async login(token: string) {
    this.token = token;
  }

  async logout() {
    this.token = undefined;
  }

  async fetchUser() {
    if (this.token) {
      const res = await getProfile(this.token);
      const user = res.data;
      runInAction(() => {
        this.user = {...user};
      });
    }
  }

  get isAuth() {
    return !!this.token;
  }
}

const rootStore = new RootStore();

reaction(
    () => rootStore.token,
    (token) => {
      token
          ? localStorage.setItem('token', token)
          : localStorage.removeItem('token');
    },
);
reaction(
    () => rootStore.token,
    async (token) => {
      if (token) await rootStore.fetchUser();
    },
);

export default rootStore;
export {RootStore, ProtectedRoute};