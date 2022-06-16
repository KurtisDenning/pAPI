import {
  Box,
  Flex,
  Text,
  Heading,
  VStack,
  Input,
  InputGroup,
  InputRightElement,
  Center,
} from "@chakra-ui/react";
import { Link } from "react-scroll";
import { React } from "react";

const Header = ({ query, setQuery }) => {
  return (
    <>
      <Box h={"90vh"} mx={50}>
        <Flex h={"80%"} alignItems={"center"} justifyContent={"center"}>
          <VStack spacing={10}>
            <Heading fontSize={["3rem", "3.5rem", "4.5rem", "5.25rem", "8rem"]}>
              The Modern
              <br /> API Platform
            </Heading>
            <Text fontSize={["1rem", "1.25rem", "1.5rem"]}>Find, connect to, & manage thousands of API's</Text>
          </VStack>
        </Flex>
        <Center>
          <InputGroup w={["80vw", "70vw", "60vw", "50vw", "40vw"]}>
            <Input
              fontSize={["1.25rem", "1.5rem"]}
              onChange={(event) => setQuery(event.target.value)}
              isInvalid
              errorBorderColor="black"
              noOfLines={1}
              placeholder="Search for an API here"
              value={query}
            ></Input>

            <InputRightElement>
              <Box
                mr={5}
                as="button"
                border="0px"
                borderRadius="2px"
                fontSize={["1.25rem", "1.5rem"]}
                fontWeight="semibold"
                bg="#ffffff"
                color="black"
              >
                <Link
                  to="APIs"
                  spy={true}
                  smooth={true}
                  offset={-50}
                  duration={500}
                >
                  &#62;
                </Link>
              </Box>
            </InputRightElement>
          </InputGroup>
        </Center>
      </Box>
    </>
  );
};

export default Header;
