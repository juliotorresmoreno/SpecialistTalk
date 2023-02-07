import React from 'react'
import Header from '../components/Header'

type LoadingPageProps = {}

const LoadingPage: React.FC<LoadingPageProps> = () => {
  const header = {
    title: 'Error ',
    description: 'programa de super poderes',
  }

  return (
    <>
      <Header {...header} />
      <main>Loading</main>
    </>
  )
}

export default LoadingPage
