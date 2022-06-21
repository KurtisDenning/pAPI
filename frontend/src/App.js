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
import theme from "./theme/index.js";
import "./theme/styles.css";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Nav from "./Components/Nav";
import Header from "./Components/Header";
import Categories from "./Components/Categories";
import APIPage from "./Components/APIPage";
import APIs from "./Components/APIs";
import Footer from "./Components/Footer";
import axios from "axios";
function App() {
  const [query, setQuery] = useState("");
  const [category, setCategory] = useState(null);
  const [APIData, setAPIData] = useState([]);
  const [categoryData, setCategoryData] = useState([]);
  const [categorysLoading, setCategorysLoading] = useState(true);
  const [APIsLoading, setAPIsLoading] = useState(true);
  const [showButton, setShowButton] = useState(false);

  // Get category data
  useEffect(() => {
    const fetchCategoryData = async () => {
      const res = await axios.get(
        "https://papi-project.herokuapp.com/api/v1/categories"
      );
      setCategoryData(res.data);
      setCategorysLoading(false);
    };
    fetchCategoryData();
  }, []);

  // Get API data
  useEffect(() => {
    const fetchAPIData = async () => {
      const res = await axios.get(
        "https://papi-project.herokuapp.com/api/v1/apidata"
      );
      setAPIData(res.data);
      setAPIsLoading(false);
    };
    fetchAPIData();
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
                  <Text textAlign={"center"} pt={20} color={"red"}>
                    This is taking longer than expected...
                  </Text>
                  <Button
                    onClick={() => {
                      window.location.reload(false);
                    }}
                  >
                    <Text textAlign={"center"}>
                      Click here to refresh the page
                    </Text>
                  </Button>
                  <Text textAlign={"center"} fontSize={"xs"} as={"i"}>
                    If the error persists, please reach out to one of our admins
                    at <b>2020000230@student.sit.ac.nz</b> .
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
          />
          <Divider />
          <APIs query={query} category={category} APIData={APIData} />
        </>
      );
    }
  }

  return (
    <ChakraProvider theme={theme}>
      <Router>
        <Nav />
        <Routes>
          <Route
            path="/pAPI"
            element={
              <>
                <Header query={query} setQuery={setQuery} />
                <Divider />
                {isLoading()}
              </>
            }
          />
          <Route exact path="/pAPI/API">
            <Route path=":id" element={<APIPage />} />
          </Route>
        </Routes>
        <Footer />
      </Router>
    </ChakraProvider>
  );
}

export default App;
