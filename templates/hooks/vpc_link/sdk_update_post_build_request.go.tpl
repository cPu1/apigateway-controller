    if err := updateVPCLinkInput(desired, latest, input, delta); err != nil {
        return nil, err
    }
