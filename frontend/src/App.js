import { React, useState } from "react";
import { ChakraProvider, Divider } from "@chakra-ui/react";
import Nav from "./Components/Nav";
import Header from "./Components/Header";
import Categories from "./Components/Categories";
import APIs from "./Components/APIs";
import Footer from "./Components/Footer";

function App() {
  const isDevelopment = true;

  const [query, setQuery] = useState("");
  const [category, setCategory] = useState(null);

  return (
    <ChakraProvider>
      <Nav />
      <Header query={query} setQuery={setQuery} />
      <Divider />
      <Categories
        setCategory={setCategory}
        category={category}
        isDevelopment={isDevelopment}
      />
      <Divider />
      <APIs query={query} category={category} isDevelopment={isDevelopment} />
      <Footer />
    </ChakraProvider>
  );
}

export default App;
