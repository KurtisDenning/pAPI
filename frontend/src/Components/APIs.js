import React from "react";
import APICard from "./APICards";
import APIData from "../Dummy-Json/API-data";
import { Center, Heading, Box } from "@chakra-ui/react";

const APIs = ({ query }) => {
  let apiCards = APIData.filter((card) => {
    if (query === "") {
      return card;
    } else if (card.name.toLowerCase().includes(query.toLowerCase())) {
      return card;
    }
  }).map((item) => <APICard key={item.id} {...item} />);

  return (
    <Box m={50}>
      <Center mb={50}>
        <Heading as={"h3"} size={"lg"} id={"APIs"}>
          API's {/*need to set state to for what results are being filtered*/}
        </Heading>
      </Center>
      {/*Need to add a search bar.*/}
      {apiCards}
    </Box>
  );
};

export default APIs;
