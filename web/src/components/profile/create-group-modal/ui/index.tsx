import { Button, Input, Modal, Text } from "@nextui-org/react";
import { useEvent, useStore } from "effector-react";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { TranslationItem, useTranslation } from "store/i18n";
import {
  $createGroupClicked,
  $groupCreated,
  $groupModalToggled,
  $isCreateGroupModalOpen,
} from "../model";

export const CreateGroupModal: React.FC = () => {
  const isOpen = useStore($isCreateGroupModalOpen);
  const groupModalToggled = useEvent($groupModalToggled);
  const createGroupClicked = useEvent($createGroupClicked);
  const [title, setTitle] = useState("");
  const navigate = useNavigate();

  const $t = useTranslation();

  useEffect(() => {
    const sub = $groupCreated.watch((id) => navigate(`/groups/${id}`));
    return () => sub.unsubscribe();
  }, []);

  return (
    <>
      <Modal
        closeButton
        aria-labelledby="modal-title"
        open={isOpen}
        onClose={() => groupModalToggled(false)}
      >
        <Modal.Header>
          <Text id="modal-title" size={18}>
            <Text b size={18}>
              {$t[TranslationItem.CreateGroupModalTitle]}
            </Text>
          </Text>
        </Modal.Header>
        <Modal.Body>
          <Input
            clearable
            bordered
            fullWidth
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            color="primary"
            size="lg"
            placeholder={$t[TranslationItem.CreateGroupModalInputText]}
          />
        </Modal.Body>
        <Modal.Footer>
          <Button
            css={{ w: "100%" }}
            size={"lg"}
            flat
            auto
            onClick={() => createGroupClicked({ title })}
          >
            {$t[TranslationItem.CreateGroupModalButtonText]}
          </Button>
        </Modal.Footer>
      </Modal>
    </>
  );
};
