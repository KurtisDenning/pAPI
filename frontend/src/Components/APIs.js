import { React, useEffect, useState } from "react";
import APICard from "./APICards";
import APIDevJSON from "../DevData/APIs.json";
import { Heading, Box, Center, Spinner } from "@chakra-ui/react";
import Pagination from "./Pagination";
import axios from "axios"

const APIs = ({ query, category, isDevelopment }) => {
  const [APIData, setAPIData] = useState([]);
  const [loading, setLoading] = useState(true);
  const [currentPage, setCurrentPage] = useState(1);
  const [itemsPerPage] = useState(3);

  useEffect(() => {
    if (isDevelopment === false) {
      const fetchAPIData = async () => {
        const res = await axios.get("https://papi-project.herokuapp.com/api/v1/apidata");
        setAPIData(res.data);
        setLoading(false);
      }
      fetchAPIData();
    } else {
      setTimeout(() => {  //Simulates loading times
        setAPIData(APIDevJSON);
        setLoading(false)
      }, 5000);
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
      .map((item) => <APICard key={item._id} {...item} loading={loading} />);

    //Get current posts
    const indexOfLastItem = currentPage * itemsPerPage;
    const indexOfFirstItem = indexOfLastItem - itemsPerPage;
    const currentItems = apiCards.slice(indexOfFirstItem, indexOfLastItem);

    //Change page
    const paginate = pageNumber => setCurrentPage(pageNumber)

    return (
      <Box>
        <Box id={"APIs"} h={1}></Box>
        <Box my={50} mx={[50, 100, 150]}>
          <Heading as={"h3"} size={"lg"} mb={"50px"}>
            API's
          </Heading>
          {currentItems}
          <Pagination itemsPerPage={itemsPerPage} totalPosts={APIData.length} paginate={paginate} />
        </Box>
      </Box>
    );
  }
};

export default APIs;
