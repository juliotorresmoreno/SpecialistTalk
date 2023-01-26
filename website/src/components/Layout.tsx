import React from 'react'
import NavBar from './NavBar'
import Container from 'react-bootstrap/Container'

type LayoutProps = {
  withChat?: boolean
} & React.PropsWithChildren

const Layout: React.FC<LayoutProps> = ({ children, withChat = false }) => {
  return (
    <>
      <NavBar />
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
