import React from "react";
import CategoryCards from "./CategoryCards";
import Data from "../Dummy-Json/Categories-data";
import { Center, Heading, Text, Box } from "@chakra-ui/react";

const Categories = ({ setCategory, category }) => {
  const categoryCards = Data.map((item) => (
    <CategoryCards
      setCategory={setCategory}
      category={category}
      key={item.id}
      {...item}
    />
  ));

  return (
    <Box m={50}>
      <Center>
        <Heading as={"h3"} size={"lg"}>
          Not sure what you are looking for?
        </Heading>
      </Center>
      <Text mt={2} mb={"80px"} color={"gray"}>
        Pick a category & start exploring!
      </Text>
      {categoryCards}
    </Box>
  );
};

export default Categories;
