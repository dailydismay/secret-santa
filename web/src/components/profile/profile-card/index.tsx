import { Button, Card, Container, Grid, Text } from "@nextui-org/react";
import { formatRelative } from "date-fns";
import { useStore } from "effector-react";
import { Link } from "react-router-dom";
import { $dateFnsLocale, TranslationItem, useTranslation } from "store/i18n";
import { Visitor } from "store/visitor/types";

export const ProfileCard: React.FC<{ visitor: Visitor }> = ({ visitor }) => {
  const $t = useTranslation();
  const dateFnsLocale = useStore($dateFnsLocale);

  return (
    <Card css={{ p: "$6", mt: "$10" }}>
      <Card.Header>
        <img
          alt="nextui logo"
          style={{ borderRadius: "10px" }}
          src={visitor.avatarURL}
          width="60px"
          height="60px"
        />
        <Grid.Container css={{ pl: "$6" }}>
          <Grid xs={12}>
            <Text h4 css={{ lineHeight: "$xs" }}>
              {visitor.firstName} {visitor.lastName}
            </Text>
          </Grid>
          <Grid xs={12}>
            <Text css={{ color: "$accents8" }}>
              VK_ID{visitor.authProviderID}
            </Text>
          </Grid>
        </Grid.Container>
      </Card.Header>
      <Card.Body css={{ py: "$2" }}>
        <Text>
          {$t[TranslationItem.AccountCreatedAt]} <br></br>
          {formatRelative(new Date(visitor.createdAt), new Date(), {
            locale: dateFnsLocale,
          })}
        </Text>
      </Card.Body>
      <Card.Footer>
        <Container css={{ p: "$0" }} justify="center">
          <Link
            isExternal
            color="primary"
            target="_blank"
            href="https://github.com/dailydismay"
          >
            {$t[TranslationItem.VisitCreatorGithub]}
          </Link>
          <Button flat css={{ w: "100%", mt: "$10" }} color="error">
            {$t[TranslationItem.NavLogout]}
          </Button>
        </Container>
      </Card.Footer>
    </Card>
  );
};
