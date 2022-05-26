import React from "react";
import { ChakraProvider, Divider } from "@chakra-ui/react";
import Nav from "./Components/Nav";
import Header from "./Components/Header";
import Categories from "./Components/Categories";
import APIs from "./Components/APIs";
import Footer from "./Components/Footer";

function App() {
  return (
    <ChakraProvider>
      <Nav />
      <Header />
      <Divider />
      <Categories />
      <Divider />
      <APIs />
      <Footer />
    </ChakraProvider>
  );
}

export default App;
