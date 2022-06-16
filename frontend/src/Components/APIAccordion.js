import { React, useState } from "react";
import axios from "axios";
import {
  Box,
  Spinner,
  Text,
  Accordion,
  AccordionIcon,
  AccordionButton,
  AccordionItem,
  AccordionPanel,
  Heading,
  Center,
} from "@chakra-ui/react";

function APIAccordion({ item, index, id, isEven }) {
  const [loading, setLoading] = useState(true);
  const [response, setResponse] = useState("");
  const [isOpen, setIsOpen] = useState(false);

  const fetchData = async () => {
    if (!isOpen) {
      const res = await axios.get(
        `https://papi-project.herokuapp.com/api/v1/apidata/${id}/${index}`
      );
      setResponse(res.data);
      setLoading(false);
      setIsOpen(true);
    }
  };

  let accordionBg;

  if (isEven) {
    accordionBg = "gray.100";
  } else {
    accordionBg = "gray.400";
  }

  return (
    <>
      <Box bg={accordionBg}>
        <Accordion allowToggle mb={"5"}>
          <AccordionItem>
            <AccordionButton onClick={fetchData}>
              <Box flex="1" textAlign="left">
                <Text fontSize={["1rem", "1.25rem"]}>
                  {item.request}
                </Text>
              </Box>
              <AccordionIcon />
            </AccordionButton>
            <AccordionPanel pb={4} onClick={fetchData} p={1}>
              <Box bg={"white"} p={5}>
                {loading && (
                  <>
                    <Center my={20}>
                      <Spinner size={"xl"} />
                    </Center>
                  </>
                )}

                {!loading && <Text>{JSON.stringify(response)}</Text>}
              </Box>
            </AccordionPanel>
          </AccordionItem>
        </Accordion>
      </Box>
    </>
  );
}

export default APIAccordion;
