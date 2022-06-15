import { React, useState, useRef } from "react";
import APICard from "./APICards";
import { Heading, Box } from "@chakra-ui/react";
import Pagination from "./Pagination";

const APIs = ({ query, category, APIData }) => {
  const [currentPage, setCurrentPage] = useState(1);
  const [itemsPerPage] = useState(3);

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
    <Box>
      <Box id={"APIs"} h={1}></Box>
      <Box my={50} mx={[50, 100, 150]}>
        <Heading as={"h3"} size={"lg"} mb={"50px"}>
          API's
        </Heading>
        {currentItems}
        <Pagination
          itemsPerPage={itemsPerPage}
          totalPosts={APIData.length}
          paginate={paginate}
        />
      </Box>
    </Box>
  );
};

export default APIs;
