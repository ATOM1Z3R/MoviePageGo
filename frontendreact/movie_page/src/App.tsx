import React from 'react'
import { Link, Outlet } from 'react-router-dom';
import './App.css'
import Footer from './components/Footer';
import HeaderBar from './components/headerBar'

class App extends React.Component {
  constructor(props: any) {
    super(props);

  }

  render(): JSX.Element{
    return (
      <div className="bg-gradient-to-tl content-center bg-fixed from-gray-600 to-gray-800">
        <HeaderBar />
        <Outlet />
        <Footer />
      </div>
      
    );
  }
}

export default App
