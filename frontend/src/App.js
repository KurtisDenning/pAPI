import { React, useState } from "react";
import { ChakraProvider, Divider } from "@chakra-ui/react";
import Nav from "./Components/Nav";
import Header from "./Components/Header";
import Categories from "./Components/Categories";
import APIs from "./Components/APIs";
import Footer from "./Components/Footer";

function App() {
  const [query, setQuery] = useState("");

  return (
    <ChakraProvider>
      <Nav />
      <Header setQuery={setQuery} />
      <Divider />
      <Categories />
      <Divider />
      <APIs query={query} />
      <Footer />
    </ChakraProvider>
  );
}

export default App;
