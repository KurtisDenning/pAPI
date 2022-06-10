import { React, useState, useEffect } from "react";
import { ChakraProvider, Divider } from "@chakra-ui/react";
import Nav from "./Components/Nav";
import Header from "./Components/Header";
import Categories from "./Components/Categories";
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
  const [categoryLoading, setCategoryLoading] = useState(true);
  const [APIsLoading, setAPIsLoading] = useState(true);

  // Get category data
  useEffect(() => {
    if (isDevelopment === false) {
      const fetchCategoryData = async () => {
        const res = await axios.get("https://papi-project.herokuapp.com/api/v1/categories")
        setCategoryData(res.data)
        setCategoryLoading(false)
      }
      fetchCategoryData()

    } else {
      setTimeout(() => {
        setCategoryData(CategoriesDevJSON);
        setCategoryLoading(false)
      }, 3000)
    }
  }, []);

  // Get API data
  useEffect(() => {
    if (isDevelopment === false) {
      const fetchAPIData = async () => {
        const res = await axios.get("https://papi-project.herokuapp.com/api/v1/apidata");
        setAPIData(res.data);
        setAPIsLoading(false);
      }
      fetchAPIData();
    } else {
      setTimeout(() => {  //Simulates loading times
        setAPIData(APIDevJSON);
        setAPIsLoading(false)
      }, 3000);
    }
  }, []);

  return (
    <ChakraProvider>
      <Nav />
      <Header query={query} setQuery={setQuery} />
      <Divider />
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
      <Footer />
    </ChakraProvider>
  );
}

export default App;
