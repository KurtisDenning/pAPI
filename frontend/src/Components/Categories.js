import { React, useState, useEffect } from "react";
import CategoryCards from "./CategoryCards";
import CategoriesDevJSON from "../DevData/Categories.json";
import { Heading, Text, Box, Center, Spinner } from "@chakra-ui/react";
import axios from "axios";

const Categories = ({ setCategory, category, isDevelopment }) => {
  const [categoryData, setCategoryData] = useState([]);
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    if (isDevelopment === false) {
      const fetchCategoryData = async () => {
        const res = await axios.get("https://papi-project.herokuapp.com/api/v1/categories")
        setCategoryData(res.data)
        setLoading(false)
      }
      fetchCategoryData()

    } else {
      setTimeout(() => {
        setCategoryData(CategoriesDevJSON);
        setLoading(false)
      }, 5000)

    }
  }, []);

  if (loading) {
    return (
      <Center>
        <Heading my={"8rem"}>
          <Spinner size={"xl"} />
        </Heading>
      </Center>
    );
  } else {

    const categoryCards = categoryData.map((item) => (
      <CategoryCards
        setCategory={setCategory}
        category={category}
        key={item._id}
        item={item}
      />
    ));

    return (
      <Box my={50} mx={[50, 100, 150]}>
        <Heading as={"h3"} size={"lg"}>
          Not sure what you are looking for?
        </Heading>
        <Text mt={4} mb={"80px"} color={"gray"}>
          Pick a category & start exploring!
        </Text>
        {categoryCards}
      </Box>
    );
  }
};

export default Categories;
