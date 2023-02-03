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
            display: 'flex',
            flex: 1,
          }}
        >
          <div></div>
        </div>
      </main>
    </>
  )
}

export default HomePage
