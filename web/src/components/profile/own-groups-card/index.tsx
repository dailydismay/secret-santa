import { Button, Card, Container, Text } from "@nextui-org/react";
import { useEvent } from "effector-react";
import { TranslationItem, useTranslation } from "store/i18n";
import { $groupModalToggled } from "../create-group-modal";
import { OwnGroupsTable } from "./own-groups-table";

export const OwnGroupsCard: React.FC = () => {
  const $t = useTranslation();
  const groupModalToggled = useEvent($groupModalToggled);

  return (
    <Card css={{ p: "$6", mt: "$10" }}>
      <Card.Header>
        <Text h4>{$t[TranslationItem.GroupsCreatedByYou]}</Text>
      </Card.Header>
      <Card.Body css={{ p: "$0" }}>
        <OwnGroupsTable></OwnGroupsTable>
      </Card.Body>
      <Card.Footer>
        <Container display="flex" justify="center">
          <Button
            size={"lg"}
            onClick={() => groupModalToggled(true)}
            color={"gradient"}
          >
            {$t[TranslationItem.CreateGroupModalButtonOpenText]}
          </Button>
        </Container>
      </Card.Footer>
    </Card>
  );
};
