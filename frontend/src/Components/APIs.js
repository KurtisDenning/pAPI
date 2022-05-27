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

const APIs = ({ setQuery, query, category }) => {
  let apiCards = APIData.filter((card) => {
    if (query === "") {
      return card;
    } else if (card.name.toLowerCase().includes(query.toLowerCase())) {
      return card;
    }
  }).map((item) => <APICard key={item.id} {...item} />);

  return (
    <Box>
      <Box id={"APIs"} h={1}></Box> {/*Remove for scooter*/}
      <Box m={50}>
        <Center mb={50}>
          <Heading as={"h3"} size={"lg"}>
            API's {/*need to set state to for what results are being filtered*/}
          </Heading>
        </Center>
        <Center>
          <InputGroup w={"80vw"} mb="50">
            <Input
              onChange={(event) => setQuery(event.target.value)}
              isInvalid
              errorBorderColor="black"
              placeholder="Search for an API here"
              value={query}
            ></Input>

            <InputRightElement>
              <Box
                as="button"
                border="0px"
                borderRadius="2px"
                fontSize="14px"
                fontWeight="semibold"
                bg="#ffffff"
                color="black"
              >
                <a href="#APIs">GO</a>
              </Box>
            </InputRightElement>
          </InputGroup>
        </Center>
        {apiCards}
      </Box>
    </Box>
  );
};

export default APIs;
