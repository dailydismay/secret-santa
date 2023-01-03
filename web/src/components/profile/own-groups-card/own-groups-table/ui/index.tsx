import {
  Badge,
  Col,
  Popover,
  Row,
  Table,
  Text,
  Tooltip,
} from "@nextui-org/react";
import { Status } from "components/shared/badges";
import { DeleteIcon, IconButton } from "components/shared/icons";
import { formatRelative } from "date-fns";
import { useGate, useStore } from "effector-react";
import { useState } from "react";
import { $dateFnsLocale, TranslationItem, useTranslation } from "store/i18n";
import { $isLoading, $listOwnGroupsGate, $ownGroups } from "../model";

export const OwnGroupsTable: React.FC = () => {
  useGate($listOwnGroupsGate);
  const $t = useTranslation();
  const ownGroups = useStore($ownGroups);
  // const isLoading = useState($isLoading);

  return (
    <Table
      shadow={false}
      striped
      headerLined
      css={{
        height: "auto",
        width: "100%",
      }}
      selectionMode="none"
    >
      <Table.Header>
        <Table.Column>{$t[TranslationItem.GroupNameTitle]}</Table.Column>
        <Table.Column>
          {$t[TranslationItem.GroupInvitationCodeTitle]}
        </Table.Column>
        <Table.Column>{$t[TranslationItem.GroupStatusTitle]}</Table.Column>
        <Table.Column hideHeader align="center">
          Controls
        </Table.Column>
      </Table.Header>
      <Table.Body>
        {ownGroups ? (
          ownGroups.map((x, key) => (
            <Table.Row key={key}>
              <Table.Cell>
                <Col>
                  <Row>
                    <Text b size={14}>
                      {x.title}
                    </Text>
                  </Row>
                </Col>
              </Table.Cell>
              <Table.Cell>
                <Row css={{ mt: "$2" }}>
                  <Popover>
                    <Popover.Trigger>
                      <Badge
                        size={"sm"}
                        onClick={() =>
                          navigator.clipboard.writeText(x.invitationCode)
                        }
                        css={{ cursor: "pointer" }}
                        isSquared
                      >
                        {x.invitationCode}
                      </Badge>
                    </Popover.Trigger>
                    <Popover.Content>
                      <Text css={{ p: "$10" }}>
                        {$t[TranslationItem.InvitationCodeCopied]}
                      </Text>
                    </Popover.Content>
                  </Popover>
                </Row>
              </Table.Cell>
              <Table.Cell>
                <Status type={x.status}>
                  {x.status === "finished"
                    ? $t[TranslationItem.GroupStatusFinished]
                    : $t[TranslationItem.GroupStatusPending]}
                </Status>
              </Table.Cell>
              <Table.Cell>
                <Tooltip
                  content={$t[TranslationItem.DeleteGroupHint]}
                  color="error"
                  onClick={() => console.log("Delete group")}
                >
                  <IconButton>
                    <DeleteIcon size={20} fill="#FF0080" />
                  </IconButton>
                </Tooltip>
              </Table.Cell>
            </Table.Row>
          ))
        ) : (
          <Table.Row key={1}>
            <Table.Cell>Loading</Table.Cell>
            <Table.Cell>Loading</Table.Cell>
            <Table.Cell>Loading</Table.Cell>
            <Table.Cell>Loading</Table.Cell>
          </Table.Row>
        )}
      </Table.Body>
    </Table>
  );
};
