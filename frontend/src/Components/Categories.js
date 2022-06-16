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
      <Heading fontSize={["1.5rem", "2rem", "3rem", "4rem"]}>
        Not sure what you are looking for?
      </Heading>
      <Text fontSize={["1rem", "1.25rem", "1.5rem"]} mt={4} mb={"80px"} color={"gray"}>
        Pick a category & start exploring!
      </Text>
      {categoryCards}
    </Box>
  );
}

export default Categories;
