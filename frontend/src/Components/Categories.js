import { React } from "react";
import CategoryCards from "./CategoryCards";
import { Heading, Text, Box } from "@chakra-ui/react";

const Categories = ({ setCategory, category, categoryData }) => {

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

export default Categories;
