import type { NextPage } from 'next';
import Player from '../components/Player';

const Home: NextPage = () => {
  return (
    <div>
      <h1>Bienvenue sur ts-migrate</h1>
      <Player />
    </div>
  );
};

export default Home;
