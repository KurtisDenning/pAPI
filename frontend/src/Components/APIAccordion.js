import { React, useState } from "react";
import axios from "axios";
import {
  Box,
  Spinner,
  Button,
  Text,
  Accordion,
  AccordionIcon,
  AccordionButton,
  AccordionItem,
  AccordionPanel,
} from "@chakra-ui/react";

function APIAccordion({ item, index, id }) {
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

  return (
    <>
      <Accordion allowToggle>
        <AccordionItem>
          <h2>
            <AccordionButton onClick={fetchData}>
              <Box flex="1" textAlign="left">
                {item.request}
              </Box>
              <AccordionIcon />
            </AccordionButton>
          </h2>
          <AccordionPanel pb={4} onClick={fetchData}>
            {loading && (
              <Box my={50} mx={[50, 100, 150]}>
                <Spinner size={"xl"} />
              </Box>
            )}

            {!loading && <Text>{JSON.stringify(response)}</Text>}
          </AccordionPanel>
        </AccordionItem>
      </Accordion>
    </>
  );
}

export default APIAccordion;
