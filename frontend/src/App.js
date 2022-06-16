import { React, useState, useEffect } from "react";
import {
  ChakraProvider,
  Divider,
  Spinner,
  Center,
  Button,
  Text,
  VStack,
} from "@chakra-ui/react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Nav from "./Components/Nav";
import Header from "./Components/Header";
import Categories from "./Components/Categories";
import APIPage from "./Components/APIPage";
import APIs from "./Components/APIs";
import Footer from "./Components/Footer";
import axios from "axios";
import CategoriesDevJSON from "./DevData/Categories.json";
import APIDevJSON from "./DevData/APIs.json";

function App() {
  const isDevelopment = true;

  const [query, setQuery] = useState("");
  const [category, setCategory] = useState(null);
  const [APIData, setAPIData] = useState([]);
  const [categoryData, setCategoryData] = useState([]);
  const [categorysLoading, setCategorysLoading] = useState(true);
  const [APIsLoading, setAPIsLoading] = useState(true);
  const [showButton, setShowButton] = useState(false);

  // Get category data
  useEffect(() => {
    if (isDevelopment === false) {
      const fetchCategoryData = async () => {
        const res = await axios.get(
          "https://papi-project.herokuapp.com/api/v1/categories"
        );
        setCategoryData(res.data);
        setCategorysLoading(false);
      };
      fetchCategoryData();
    } else {
      setTimeout(() => {
        // Simulates loading times
        setCategoryData(CategoriesDevJSON);
        setCategorysLoading(false);
      }, 3000);
    }
  }, []);

  // Get API data
  useEffect(() => {
    if (isDevelopment === false) {
      const fetchAPIData = async () => {
        const res = await axios.get(
          "https://papi-project.herokuapp.com/api/v1/apidata"
        );
        setAPIData(res.data);
        setAPIsLoading(false);
      };
      fetchAPIData();
    } else {
      setTimeout(() => {
        // Simulates loading times
        setAPIData(APIDevJSON);
        setAPIsLoading(false);
      }, 3000);
    }
  }, []);

  useEffect(() => {
    setTimeout(() => {
      setShowButton(true);
    }, 10000);
  }, []);

  function isLoading() {
    if (APIsLoading || categorysLoading) {
      return (
        <>
          <Center my={40}>
            <VStack mx={50}>
              <Spinner size={"xl"} />

              {showButton && (
                <>
                  <Text pt={20} color={"red"}>
                    This is taking longer than expected...
                  </Text>
                  <Button
                    onClick={() => {
                      window.location.reload(false);
                    }}
                  >
                    <Text>Click here to refresh the page</Text>
                  </Button>
                  <Text fontSize={"xs"} as={"i"}>
                    If the error persists, please reach out to one of our admins
                    at <b>testing@blabla.com</b> .
                  </Text>
                </>
              )}
            </VStack>
          </Center>
        </>
      );
    } else {
      return (
        <>
          <Categories
            setCategory={setCategory}
            category={category}
            categoryData={categoryData}
            isDevelopment={isDevelopment}
          />
          <Divider />
          <APIs
            query={query}
            category={category}
            APIData={APIData}
            isDevelopment={isDevelopment}
          />
        </>
      );
    }
  }

  return (
    <ChakraProvider>
      <Nav />
      <Router>
        <Routes>
          <Route
            path="/"
            element={
              <>
                <Header query={query} setQuery={setQuery} />
                <Divider />
                {isLoading()}
              </>
            }
          />
          <Route path="/api">
            <Route path=":id" element={<APIPage />} />
          </Route>
        </Routes>
      </Router>
      <Footer />
    </ChakraProvider>
  );
}

export default App;
