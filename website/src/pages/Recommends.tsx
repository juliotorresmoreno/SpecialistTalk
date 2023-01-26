import React from 'react'
import Header from '../components/Header'

const RecommendsPage: React.FC = () => {
  const header = {
    title: 'Recommends',
    description: 'programa de super poderes',
  }

  return (
    <>
      <Header {...header} />
      <main>Recommends</main>
    </>
  )
}

export default RecommendsPage
