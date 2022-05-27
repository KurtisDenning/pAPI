import React from "react";
import APICard from "./APICards";
import APIData from "../Dummy-Json/API-data";
import {
  Center,
  Heading,
  Box,
  Input,
  InputGroup,
  InputRightElement,
} from "@chakra-ui/react";

const APIs = ({ query, category }) => {
  let apiCat = APIData.filter((card) => {
    if (category === "") {
      return card;
    } else if (card.category.toLowerCase() === category.toLowerCase()) {
      return card;
    }
  });

  let apiCards = apiCat
    .filter((card) => {
      if (query === "") {
        return card;
      } else if (card.name.toLowerCase().includes(query.toLowerCase())) {
        return card;
      }
    })
    .map((item) => <APICard key={item.id} {...item} />);

  return (
    <Box>
      <Box id={"APIs"} h={1}></Box> {/*Remove for scooter*/}
      <Box m={50}>
        <Center mb={50}>
          <Heading as={"h3"} size={"lg"}>
            API's {/*need to set state to for what results are being filtered*/}
          </Heading>
        </Center>
        {apiCards}
      </Box>
    </Box>
  );
};

export default APIs;
