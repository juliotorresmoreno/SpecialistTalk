
import React from 'react';
import Header from '../components/Header';

const NotFoundPage: React.FC = () => {
  const header = {
    title: 'NotFound',
    description: 'programa de super poderes'
  };

  return (
    <>
      <Header {...header} />
      <main>
        NotFound
      </main>
    </>
  );
}

export default NotFoundPage;
