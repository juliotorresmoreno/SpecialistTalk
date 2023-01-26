import React from 'react'
import Header from '../components/Header'

const HomePage: React.FC = () => {
  const header = {
    title: 'Home',
    description: 'programa de super poderes',
  }
  return (
    <>
      <Header {...header} />
      <main>
        <div
          style={{
            display: 'grid',
            gridTemplateRows: 'repeat(3, 100px)',
            gridTemplateColumns: 'repeat(3, 1fr)',
          }}
        >
          <div
            style={{
              gridArea: '2 / 1 / 4 / 4',
              background: 'red',
            }}
          ></div>
        </div>
      </main>
    </>
  )
}

export default HomePage
