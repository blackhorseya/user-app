import {observer} from 'mobx-react-lite';
import rootStore from '@/store';

const Profile = () => {
  return (
      <div className="card md:card-side card-bordered">
        <figure>
          <img src={rootStore.user?.picture}/>
        </figure>
        <div className="p-2 card-content">
          <div className="card-title">
            {rootStore.user?.name} (
            <span className="capitalize">{rootStore.user?.name}</span>)
          </div>
          <a href={`mailto:${rootStore.user?.email}`} className="link">
            {rootStore.user?.email}
          </a>
        </div>
      </div>
  );
};

export default observer(Profile);