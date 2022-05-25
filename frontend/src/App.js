import React from 'react';
import { ChakraProvider } from '@chakra-ui/react';
import Nav from './Components/Nav';
import Header from './Components/Header';

function App() {
  return (
    <ChakraProvider>
      <Nav/>
      <Header/>
    </ChakraProvider>
  );
}

export default App;
