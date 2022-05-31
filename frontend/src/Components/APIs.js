import { React, useEffect, useState } from "react";
import APICard from "./APICards";
import { Heading, Box } from "@chakra-ui/react";

const APIs = ({ query, category }) => {
  const [APIData, setAPIData] = useState([]);

  useEffect(() => {
    fetch("https://papi-project.herokuapp.com/api/v1/apidata")
      .then((res) => res.json())
      .then((data) => setAPIData(data));
  }, []);

  console.log(APIData);

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
      } else if (card.name.toLowerCase().includes(query.toLowerCase())) {
        return card;
      } else {
        return null;
      }
    })
    .map((item) => <APICard key={item._id} {...item} />);

  return (
    <Box>
      <Box id={"APIs"} h={1}></Box> {/*Remove for scooter*/}
      <Box my={50} mx={[50, 100, 150]}>
        <Heading as={"h3"} size={"lg"} mb={"50px"}>
          API's
        </Heading>
        {apiCards}
      </Box>
    </Box>
  );
};

export default APIs;
