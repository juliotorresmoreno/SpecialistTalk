import React, { Fragment } from 'react'
import Header from '../components/Header'
import styl from 'styled-components'

const Container = styl.div`
  display: flex;
  flex-direction: column;
  flex: 1;
  border-right: 1px solid var(--bs-gray-200);
`

const Menu = styl.div`
  border-right: 1px solid var(--bs-gray-200);
`

const ForumPage: React.FC = () => {
  const header = {
    title: 'Foro',
    description: 'programa de super poderes',
  }

  const links = new Array(10).fill(['#', 'Link'])

  return (
    <div className="page">
      <Menu>
        <nav className="nav nav-tabs">
          <a className="nav-link active" href="#">
            Active
          </a>
          {links.map(([link, title], index) => (
            <a key={'link-' + index} className="nav-link" href={link}>
              {title}
            </a>
          ))}
          <a className="nav-link" href="#">
            Link
          </a>
          <a className="nav-link disabled" href="#">
            Disabled
          </a>
        </nav>
      </Menu>
      <Container>
        <Header {...header} />
        <main>
          <section>Home</section>
        </main>
      </Container>
    </div>
  )
}

export default ForumPage
