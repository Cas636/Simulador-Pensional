import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import './assets/index.css';
import Header from './components/Header';
import Footer from './components/Footer';
import Menu from './components/Menu';
import Home from './pages/Home';
import PensionReformPage from './pages/PensionReformPage';



const App = () => {
  return (
    <div>
    <div>
        <Header/>
    </div>
    <Router>
      <div>
        <Menu />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/PensionReformPage" element={<PensionReformPage />} />
        </Routes>
      </div>
    </Router>
    <div>
        <Footer/>
    </div>
    </div>
  );
};

export default App;