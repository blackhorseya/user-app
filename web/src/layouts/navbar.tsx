import {Link} from 'react-router-dom';
import {observer} from 'mobx-react-lite';
import rootStore from '@/store';
import {getLoginUrl} from '@/utils/api';

const loginUrl = decodeURIComponent(await getLoginUrl());

const Navbar = () => {
  return (
      <div className="items-center navbar bg-primary text-primary-content">
        <div className="navbar-start">
          <Link className="text-2xl btn btn-ghost" to="">
            Sean Side
          </Link>
        </div>
        <div className="space-x-2 navbar-end">
          {!rootStore.isAuth ? (
              <button
                  className="btn "
                  onClick={() => {
                    window.location.href = loginUrl;
                  }}
              >
                Login
              </button>
          ) : (
              <div className="flex items-center space-x-2">
                <div className="dropdown dropdown-end">
                  <div tabIndex={0} className="cursor-pointer avatar">
                    <div className="w-12 h-12 rounded-full">
                      {rootStore.user && (
                          <img
                              src={rootStore.user.picture}
                              alt={rootStore.user.name}
                          />
                      )}
                    </div>
                  </div>
                  <ul
                      tabIndex={0}
                      className="p-2 shadow menu dropdown-content bg-base-100 text-base-content rounded-box"
                  >
                    <li>
                      <Link to="/profile">Profile</Link>
                    </li>
                    <li>
                      <a
                          className="text-red-500"
                          onClick={(e) => {
                            rootStore.logout();
                          }}
                      >
                        Logout
                      </a>
                    </li>
                  </ul>
                </div>
              </div>
          )}
        </div>
      </div>
  );
};

export default observer(Navbar);