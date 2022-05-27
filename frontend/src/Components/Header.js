import {
  Box,
  Flex,
  Text,
  Heading,
  VStack,
  Input,
  InputGroup,
  InputRightElement,
  Button,
} from "@chakra-ui/react";
import { React } from "react";

const Header = ({ setQuery }) => {
  return (
    <>
      <Box h={"92vh"} mx={50}>
        <Flex h={"80%"} alignItems={"center"} justifyContent={"center"}>
          <VStack spacing={10}>
            <Heading as={"h1"} size={"3xl"}>
              The Modern
              <br /> API Platform
            </Heading>
            <Text>Find, connect to, & manage thousands of APIs</Text>
          </VStack>
        </Flex>
        <Flex justifyContent={"center"}>
          <InputGroup w={"80vw"}>
            <Input
              onChange={(event) => setQuery(event.target.value)}
              isInvalid
              errorBorderColor="black"
              placeholder="Search for an API here"
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
        </Flex>
      </Box>
    </>
  );
};

export default Header;
