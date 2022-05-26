import React from "react";
import { ChakraProvider, Divider } from "@chakra-ui/react";
import Nav from "./Components/Nav";
import Header from "./Components/Header";
import Categories from "./Components/Categories";

function App() {
  return (
    <ChakraProvider>
      <Nav />
      <Header />
      <Divider />
      <Categories />
      <Divider />
    </ChakraProvider>
  );
}

export default App;
