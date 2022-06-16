import { React, useState } from "react";
import APICard from "./APICards";
import { Heading, Box, SimpleGrid } from "@chakra-ui/react";
import Pagination from "./Pagination";

const APIs = ({ query, category, APIData }) => {
  const [currentPage, setCurrentPage] = useState(1);
  const [itemsPerPage] = useState(4);

  let apiCat = APIData.filter((apiData) => {
    if (category === null) {
      return apiData;
    } else if (apiData.categories.includes(category._id)) {
      return apiData;
    } else {
      return null;
    }
  });

  let apiCards = apiCat
    .filter((card) => {
      if (query === "") {
        return card;
      } else if (card.title.toLowerCase().includes(query.toLowerCase())) {
        return card;
      } else {
        return null;
      }
    })
    .map((item) => <APICard key={item._id} {...item} />);

  //Get current posts
  const indexOfLastItem = currentPage * itemsPerPage;
  const indexOfFirstItem = indexOfLastItem - itemsPerPage;
  const currentItems = apiCards.slice(indexOfFirstItem, indexOfLastItem);

  //Change page
  const paginate = (pageNumber) => setCurrentPage(pageNumber);

  return (
    <Box id={"APIs"} my={50} mx={[50, 100, 150]}>
      <Heading fontSize={["1.5rem", "2rem", "3rem", "4rem"]} mb={"50px"}>
        API's
      </Heading>

      <SimpleGrid
        maxW={"100vw"}
        columns={[1, null, null, 2]}
      >
        {currentItems}
      </SimpleGrid>
      <Pagination
        itemsPerPage={itemsPerPage}
        totalPosts={APIData.length}
        paginate={paginate}
      />
    </Box>
  );
};

export default APIs;
