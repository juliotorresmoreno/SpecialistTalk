import React from 'react'
import NavBar from './NavBar'
import GuessNavBar from './GuessNavBar'
import Container from 'react-bootstrap/Container'
import { useAppSelector } from '../store/hooks'

type LayoutProps = {
  withChat?: boolean
} & React.PropsWithChildren

const Layout: React.FC<LayoutProps> = ({ children, withChat = false }) => {
  const session = useAppSelector((state) => state.auth.session)

  return (
    <>
      {session ? <GuessNavBar /> : <NavBar />}

      <Container id="container">
        <div className="content">
          <div>{children}</div>
          <aside>aside</aside>
        </div>
        {withChat ? (
          <div id="chats">
            <div>Chats</div>
          </div>
        ) : null}
      </Container>
    </>
  )
}

export default Layout
